public class CarsAssemble {

    public double productionRatePerHour(int speed) {
        final int carsPerHour = 221;
        if (1 <= speed && speed <= 4) {
            return 1.0 * carsPerHour * speed;
        } else if (5 <= speed && speed <= 8) {
            return 0.9 * carsPerHour * speed;
        } else if (speed == 9) {
            return 0.8 * carsPerHour * speed;
        } else {
            return 0.77 * carsPerHour * speed;
        }
    }

    public int workingItemsPerMinute(int speed) {
        final double carsPerMinute = 221 / 60.0;
        if (1 <= speed && speed <= 4) {
            return (int) (1.0 * carsPerMinute * speed);
        } else if (5 <= speed && speed <= 8) {
            return (int) (0.9 * carsPerMinute * speed);
        } else if (speed == 9) {
            return (int) (0.8 * carsPerMinute * speed);
        } else {
            return (int) (0.77 * carsPerMinute * speed);
        }
    }
}
