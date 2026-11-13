class ReverseString {

    String reverse(String inputString) {
        StringBuilder output = new StringBuilder();
        char[] inputChars = inputString.toCharArray();
        for (int i = inputString.length() - 1; i >= 0; i--) {
            output.append(inputChars[i]);
        }
        return output.toString();
    }
}
