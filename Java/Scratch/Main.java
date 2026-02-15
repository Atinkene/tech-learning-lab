public class Main {
    public static void main(String[] args) {
    Classe classe = new Classe("Math");
    Student student1 = new Student("Alice", 20, "alice@example.com");
    Student student2 = new Student("Bob", 22, "bob@example.com");
    classe.addStudent(student1);
    classe.addStudent(student2);
    System.out.println("Class: " + classe.getName());
    for (Student student : classe.getStudents()) {
        System.out.println("Student: " + student.getFirstname() + ", Age: " + student.getAge() + ", Email: " + student.getEmail());
    }
}
}
