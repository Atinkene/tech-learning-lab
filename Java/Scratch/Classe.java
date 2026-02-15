import java.util.*;
import java.util.List;

public class Classe {
    private String name;
    private List<Student> students;

    public Classe(String name) {
        this.name = name;
        this.students = new ArrayList<>();
    }

    public void addStudent(Student student) {
        students.add(student);
    }

    public String getName() {
        return name;
    }

    public List<Student> getStudents() {
        return students;
    }
}

