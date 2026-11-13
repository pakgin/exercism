import java.util.stream.Collectors;
import java.util.Arrays;
import java.util.List;

/*

Since this exercise has a difficulty of > 4 it doesn't come
with any starter implementation.
This is so that you get to practice creating classes and methods
which is an important part of programming in Java.

Please remove this comment when submitting your solution.

*/

public class Anagram {
    private String word;

    public Anagram(String word) {
        this.word = word;
    }

    public List<String> match(List<String> list) {
        return list.stream()
            .filter(w -> Arrays.equals(this.word.toLowerCase().chars().sorted().toArray(),
                    w.toLowerCase().chars().sorted().toArray())
                    && !w.toLowerCase().equals(this.word.toLowerCase()))
            .collect(Collectors.toList());
    }
}