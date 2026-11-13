class IsogramChecker {

    boolean isIsogram(String phrase) {
        char chars[] = phrase.toLowerCase().toCharArray();
        boolean counts[] = new boolean[26];

        for (int i = 0;i < counts.length;i++) {
            counts[i] = false;
        }


        for (int i = 0;i < chars.length;i++) {
            int idx = chars[i] - 'a';

            if (0 > idx || idx > 25) { continue; }

            else {
                if (counts[idx]) {
                    return false;
                }

                counts[idx] = true;
            }
        }

        return true;
    }

}
