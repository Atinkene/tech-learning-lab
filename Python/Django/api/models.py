from django.db import models
from django.contrib.auth.models import AbstractUser
from django.core.validators import MinValueValidator, MaxValueValidator
from django.utils import timezone
import uuid

STATUT_CHOICES = [
    ('actif', 'Actif'),
    ('inactif', 'Inactif'),
    ('suspendu', 'Suspendu'),
    ('annule', 'Annulé'),
    ('reussi', 'Réussi'),
    ('echoue', 'Échoué'),
    ('en_attente', 'En attente'),
    ('lu', 'Lu'),
    ('non_lu', 'Non lu'),
]

class BaseModel(models.Model):
    """Base model with common fields"""
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    
    class Meta:
        abstract = True

# Consider using AbstractUser for better Django auth integration
class Utilisateur(BaseModel):
    prenom = models.CharField(max_length=100, verbose_name="Prénom")
    nom = models.CharField(max_length=100, verbose_name="Nom")
    email = models.EmailField(unique=True)
    telephone = models.CharField(
        max_length=15, 
        blank=True, 
        null=True, 
        unique=True,
        help_text="Format: +221123456789"
    )
    nom_utilisateur = models.CharField(max_length=50, unique=True)
    mot_de_passe = models.CharField(max_length=128)
    
    ROLE_CHOICES = [
        ('admin', 'Admin'),
        ('utilisateur', 'Utilisateur'),
        ('annonceur', 'Annonceur'),
    ]
    role = models.CharField(max_length=20, choices=ROLE_CHOICES, default='utilisateur')
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    
    # Additional useful fields
    date_derniere_connexion = models.DateTimeField(blank=True, null=True)
    photo_profil = models.URLField(blank=True, null=True)
    bio = models.TextField(max_length=500, blank=True)
    
    class Meta:
        verbose_name = "Utilisateur"
        verbose_name_plural = "Utilisateurs"
        indexes = [
            models.Index(fields=['email']),
            models.Index(fields=['nom_utilisateur']),
        ]

    def __str__(self):
        return f"{self.prenom} {self.nom}"

class Categorie(BaseModel):
    nom = models.CharField(max_length=100, unique=True)
    description = models.TextField(blank=True, null=True)
    parent = models.ForeignKey(
        'self', 
        on_delete=models.CASCADE, 
        blank=True, 
        null=True, 
        related_name='sous_categories'
    )
    icone = models.URLField(blank=True, null=True)
    ordre = models.PositiveIntegerField(default=0)
    
    class Meta:
        verbose_name = "Catégorie"
        verbose_name_plural = "Catégories"
        ordering = ['ordre', 'nom']

    def __str__(self):
        return self.nom

class Annonce(BaseModel):
    titre = models.CharField(max_length=200)
    description = models.TextField()
    date_publication = models.DateTimeField(auto_now_add=True)
    date_expiration = models.DateTimeField()
    utilisateur = models.ForeignKey(
        Utilisateur, 
        on_delete=models.CASCADE, 
        related_name='annonces'
    )
    
    # Fix: Use consistent status choices
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    
    TYPE_CHOICES = [
        ('vente', 'Vente'),
        ('location', 'Location'),
        ('service', 'Service'),
        ('emploi', 'Emploi'),
    ]
    type = models.CharField(max_length=10, choices=TYPE_CHOICES, default='vente')
    categorie = models.ForeignKey(
        Categorie, 
        on_delete=models.SET_NULL, 
        null=True, 
        related_name='annonces'
    )
    
    # Additional useful fields
    prix = models.DecimalField(
        max_digits=12, 
        decimal_places=2, 
        blank=True, 
        null=True,
        validators=[MinValueValidator(0)]
    )
    prix_negociable = models.BooleanField(default=False)
    vues = models.PositiveIntegerField(default=0)
    est_premium = models.BooleanField(default=False)
    tags = models.CharField(max_length=500, blank=True, help_text="Tags séparés par des virgules")
    
    class Meta:
        verbose_name = "Annonce"
        verbose_name_plural = "Annonces"
        ordering = ['-date_publication']
        indexes = [
            models.Index(fields=['statut', 'date_publication']),
            models.Index(fields=['categorie', 'type']),
            models.Index(fields=['utilisateur']),
        ]

    def __str__(self):
        return self.titre
    
    def is_expired(self):
        return timezone.now() > self.date_expiration

class Media(BaseModel):
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, related_name='medias')
    description = models.CharField(max_length=255, blank=True, null=True)
    url = models.URLField()
    
    TYPE_CHOICES = [
        ('image', 'Image'),
        ('audio', 'Audio'),
        ('video', 'Vidéo'),
        ('document', 'Document'),
    ]
    type = models.CharField(max_length=10, choices=TYPE_CHOICES, default='image')
    ordre = models.PositiveIntegerField(default=0)
    taille_fichier = models.PositiveIntegerField(blank=True, null=True, help_text="Taille en bytes")
    
    class Meta:
        verbose_name = "Média"
        verbose_name_plural = "Médias"
        ordering = ['ordre']

    def __str__(self):
        return f"{self.type} for {self.annonce.titre}"

class Abonnement(BaseModel):
    nom = models.CharField(max_length=100)
    duree = models.IntegerField(help_text="Durée en mois", validators=[MinValueValidator(1)])
    prix = models.DecimalField(max_digits=10, decimal_places=2, validators=[MinValueValidator(0)])
    
    TYPE_CHOICES = [
        ('gratuit', 'Gratuit'),
        ('premium', 'Premium'),
        ('professionnel', 'Professionnel'),
    ]
    type = models.CharField(max_length=15, choices=TYPE_CHOICES, default='gratuit')
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    
    # Features
    nb_annonces_max = models.PositiveIntegerField(default=5)
    nb_photos_max = models.PositiveIntegerField(default=3)
    support_prioritaire = models.BooleanField(default=False)
    mise_en_avant = models.BooleanField(default=False)
    
    class Meta:
        verbose_name = "Abonnement"
        verbose_name_plural = "Abonnements"

    def __str__(self):
        return f"{self.nom} ({self.duree} mois)"

class UtilisateurAbonnement(BaseModel):
    utilisateur = models.ForeignKey(
        Utilisateur, 
        on_delete=models.CASCADE, 
        related_name='abonnements'
    )
    abonnement = models.ForeignKey(
        Abonnement, 
        on_delete=models.CASCADE, 
        related_name='utilisateur_abonnements'
    )
    date_souscription = models.DateTimeField(auto_now_add=True)
    date_expiration = models.DateTimeField()
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    auto_renouvellement = models.BooleanField(default=False)
    
    class Meta:
        verbose_name = "Abonnement Utilisateur"
        verbose_name_plural = "Abonnements Utilisateurs"

    def __str__(self):
        return f"{self.utilisateur.nom_utilisateur} - {self.abonnement.nom}"
    
    def is_active(self):
        return self.statut == 'actif' and timezone.now() < self.date_expiration

class Paiement(BaseModel):
    utilisateur_abonnement = models.ForeignKey(
        UtilisateurAbonnement, 
        on_delete=models.CASCADE, 
        related_name='paiements'
    )
    date_paiement = models.DateTimeField(auto_now_add=True)
    montant = models.DecimalField(max_digits=10, decimal_places=2, validators=[MinValueValidator(0)])
    
    MODE_CHOICES = [
        ('carte', 'Carte'),
        ('paypal', 'PayPal'),
        ('virement', 'Virement'),
        ('mobile_money', 'Mobile Money'),
        ('orange_money', 'Orange Money'),
        ('wave', 'Wave'),
    ]
    mode = models.CharField(max_length=15, choices=MODE_CHOICES, default='carte')
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='en_attente')
    
    # Transaction details
    reference_transaction = models.CharField(max_length=100, unique=True)
    details_transaction = models.JSONField(blank=True, null=True)
    
    class Meta:
        verbose_name = "Paiement"
        verbose_name_plural = "Paiements"

    def __str__(self):
        return f"Paiement de {self.montant} FCFA par {self.utilisateur_abonnement.utilisateur.nom_utilisateur}"

class Conversation(BaseModel):
    """Model for conversation threads"""
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, related_name='conversations')
    participants = models.ManyToManyField(Utilisateur, related_name='conversations')
    derniere_activite = models.DateTimeField(auto_now=True)
    
    class Meta:
        verbose_name = "Conversation"
        verbose_name_plural = "Conversations"

class Message(BaseModel):
    conversation = models.ForeignKey(Conversation, on_delete=models.CASCADE, related_name='messages')
    expediteur = models.ForeignKey(Utilisateur, on_delete=models.CASCADE, related_name='messages_envoyes')
    contenu = models.TextField()
    date_envoi = models.DateTimeField(auto_now_add=True)
    date_lecture = models.DateTimeField(blank=True, null=True)
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    
    class Meta:
        verbose_name = "Message"
        verbose_name_plural = "Messages"
        ordering = ['-date_envoi']

    def __str__(self):
        return f"Message from {self.expediteur.nom_utilisateur}"

class Favori(BaseModel):
    utilisateur = models.ForeignKey(Utilisateur, on_delete=models.CASCADE, related_name='favoris')
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, related_name='favoris')
    date_ajout = models.DateTimeField(auto_now_add=True)

    class Meta:
        unique_together = ('utilisateur', 'annonce')
        verbose_name = "Favori"
        verbose_name_plural = "Favoris"

    def __str__(self):
        return f"{self.utilisateur.nom_utilisateur} ♥ {self.annonce.titre}"

class Commentaire(BaseModel):
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, related_name='commentaires')
    utilisateur = models.ForeignKey(Utilisateur, on_delete=models.CASCADE, related_name='commentaires')
    note = models.IntegerField(validators=[MinValueValidator(1), MaxValueValidator(5)])
    contenu = models.TextField()
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='actif')
    
    class Meta:
        verbose_name = "Commentaire"
        verbose_name_plural = "Commentaires"
        unique_together = ('utilisateur', 'annonce')  # Un commentaire par utilisateur par annonce

    def __str__(self):
        return f"Commentaire de {self.utilisateur.nom_utilisateur} sur {self.annonce.titre}"

class Localisation(BaseModel):
    annonce = models.OneToOneField(Annonce, on_delete=models.CASCADE, related_name='localisation')
    adresse = models.CharField(max_length=255)
    ville = models.CharField(max_length=100)
    region = models.CharField(max_length=100)
    pays = models.CharField(max_length=100, default="Sénégal")
    code_postal = models.CharField(max_length=20, blank=True)
    latitude = models.DecimalField(max_digits=9, decimal_places=6, blank=True, null=True)
    longitude = models.DecimalField(max_digits=9, decimal_places=6, blank=True, null=True)
    
    class Meta:
        verbose_name = "Localisation"
        verbose_name_plural = "Localisations"
        indexes = [
            models.Index(fields=['ville', 'region']),
        ]

    def __str__(self):
        return f"{self.ville}, {self.region} - {self.annonce.titre}"

# Additional models you might want to consider

class Signalement(BaseModel):
    """Model for reporting inappropriate content"""
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, related_name='signalements')
    utilisateur = models.ForeignKey(Utilisateur, on_delete=models.CASCADE, related_name='signalements')
    
    MOTIF_CHOICES = [
        ('spam', 'Spam'),
        ('contenu_inapproprie', 'Contenu inapproprié'),
        ('arnaque', 'Arnaque'),
        ('fausse_annonce', 'Fausse annonce'),
        ('autre', 'Autre'),
    ]
    motif = models.CharField(max_length=20, choices=MOTIF_CHOICES)
    description = models.TextField()
    statut = models.CharField(max_length=20, choices=STATUT_CHOICES, default='en_attente')
    
    class Meta:
        unique_together = ('utilisateur', 'annonce')
        verbose_name = "Signalement"
        verbose_name_plural = "Signalements"

class Notification(BaseModel):
    """Model for user notifications"""
    utilisateur = models.ForeignKey(Utilisateur, on_delete=models.CASCADE, related_name='notifications')
    titre = models.CharField(max_length=200)
    contenu = models.TextField()
    
    TYPE_CHOICES = [
        ('message', 'Nouveau message'),
        ('favori', 'Annonce ajoutée aux favoris'),
        ('commentaire', 'Nouveau commentaire'),
        ('expiration', 'Annonce expire bientôt'),
        ('systeme', 'Notification système'),
    ]
    type = models.CharField(max_length=15, choices=TYPE_CHOICES)
    lue = models.BooleanField(default=False)
    date_lecture = models.DateTimeField(blank=True, null=True)
    
    # Optional reference to related objects
    annonce = models.ForeignKey(Annonce, on_delete=models.CASCADE, blank=True, null=True)
    
    class Meta:
        verbose_name = "Notification"
        verbose_name_plural = "Notifications"
        ordering = ['-created_at']