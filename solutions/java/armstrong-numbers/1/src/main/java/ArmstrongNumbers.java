class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
        String stringNumber = Integer.toString(numberToCheck);
        int numberOfDigits = stringNumber.length();

        int sum = 0;
        for (char ch : stringNumber.toCharArray()) {
            int digit = ch - '0';
            int term = (int) Math.pow(digit, numberOfDigits);
            sum += term;
        }
        return sum == numberToCheck;
    }
}
