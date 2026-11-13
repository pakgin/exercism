class SqueakyClean {

    static String clean(String identifier) {
        StringBuilder output = new StringBuilder();
        boolean capitalizeNext = false;

        for (char ch : identifier.toCharArray()) {
            if (Character.isWhitespace(ch)) {
                output.append('_');
            } else if (ch == '-') {
                capitalizeNext = true;
            } else {
                char letter = fromLeetspeak(ch);
                if (Character.isLetter(letter)) {
                    output.append(
                        capitalizeNext ? Character.toUpperCase(letter) : letter
                    );
                    capitalizeNext = false;
                }
            }
        }
        return output.toString();
    }

    private static char fromLeetspeak(char ch) {
        switch (ch) {
            case '4':
                return 'a';
            case '3':
                return 'e';
            case '0':
                return 'o';
            case '1':
                return 'l';
            case '7':
                return 't';
            default:
                return ch;
        }
    }
}
