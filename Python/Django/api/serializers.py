from rest_framework import serializers
from .models import Utilisateur, Annonce, Categorie

class UtilisateurSerializer(serializers.ModelSerializer):
    class Meta:
        model = Utilisateur
        fields = ['id', 'nom', 'prenom', 'email', 'telephone', 'adresse', 'role']

class CategorieSerializer(serializers.ModelSerializer):
    class Meta:
        model = Categorie
        fields = '__all__'

class AnnonceSerializer(serializers.ModelSerializer):
    utilisateur = UtilisateurSerializer(read_only=True)

    class Meta:
        model = Annonce
        fields = '__all__'
