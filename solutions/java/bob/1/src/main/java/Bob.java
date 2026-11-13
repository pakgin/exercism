class Bob {

    String hey(String input) {
        String trimmed = input.strip();
        boolean isSilent = trimmed.isEmpty();
        boolean isQuestion = trimmed.endsWith("?");
        boolean isYelling =
            trimmed.chars().anyMatch(Character::isLetter) &&
            trimmed.toUpperCase().equals(trimmed);

        if (isSilent) {
            return "Fine. Be that way!";
        } else if (isYelling && isQuestion) {
            return "Calm down, I know what I'm doing!";
        } else if (isYelling) {
            return "Whoa, chill out!";
        } else if (isQuestion) {
            return "Sure.";
        } else {
            return "Whatever.";
        }
    }
}
