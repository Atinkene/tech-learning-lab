
public class Student {
    private String firstname;
    private int age;
    private String email;

    public Student(String firstname, int age, String email) {
        this.firstname = firstname;
        this.age = age;
        this.email = email;
    }

    public String getFirstname() {
        return firstname;
    }

    public int getAge() {
        return age;
    }

    public String getEmail() {
        return email;
    }
}
