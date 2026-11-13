class RaindropConverter {

    String convert(int number) {
        String ret = "";

        if (number % 3 == 0)
            ret = ret + "Pling";

        if (number % 5 == 0)
            ret = ret + "Plang";

        if (number % 7 == 0)
            ret = ret + "Plong";

        if (ret == "")
            ret = ret + Integer.toString(number);

        return ret;
    }

}
