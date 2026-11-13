import java.util.ArrayList;
import java.util.Arrays;

/*

Since this exercise has a difficulty of > 4 it doesn't come
with any starter implementation.
This is so that you get to practice creating classes and methods
which is an important part of programming in Java.

Please remove this comment when submitting your solution.

*/

public class Anagram {
    String _word;

    public Anagram(String word) {
        _word = word;
    }

    public ArrayList<String> match(ArrayList<String> list) {
        ArrayList ret = new ArrayList<String>();

        for (int i = 0;i < list.size();i++) {
            if (list.get(i).length() != _word.length())
                continue;

            char a[] = _word.toCharArray();
            char b[] = list.get(i).toCharArray();
            Arrays.sort(a);
            Arrays.sort(b);

            for (int j = 0;j < a.length;j++) {
                if (a[j] != b[j])
                    continue;
            }

            ret.add(list.get(i));
        }

        return ret;
    }
}