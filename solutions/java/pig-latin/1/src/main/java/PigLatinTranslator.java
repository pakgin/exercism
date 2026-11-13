public class PigLatinTranslator {
    boolean isVowel(char c) {
        switch (c) {
            case 'a':
            case 'A':
            case 'e':
            case 'E':
            case 'i':
            case 'I':
            case 'o':
            case 'O':
            case 'u':
            case 'U':
                return true;
            default:
                return false;
        }
    }

    boolean isConsontant(char c) {
        return !isVowel(c);
    }

    public String translate(String word) {
        int size = word.length();
        String pigL = "";
        int i = 0;
        if (!word.startsWith("yt") && !word.startsWith("xr")) {
            for (i = 0; i < size; i++) {
                if (isVowel(word.charAt(i))) {
                    break;
                } else if (i < size - 1 && word.charAt(i) == 'q' && word.charAt(i + 1) == 'u') {
                    i = i + 2;
                    break;
                } else if (word.charAt(i) == 'y' && ((i > 0 && isConsontant(word.charAt(i - 1))) || i == size - 1)) {
                    break;
                }
            }
        }

        if (i == size) {
            return word;
        }

        pigL = word.substring(i);
        pigL = pigL + word.substring(0, i);
        pigL = pigL + "ay";

        return pigL;
    }
}
