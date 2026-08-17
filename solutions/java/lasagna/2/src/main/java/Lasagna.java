public class Lasagna {
    private static final int EXPECTED_MINUTES = 40;
    private static final int LAYER_PREP_TIME = 2;
    public int expectedMinutesInOven() {
        return EXPECTED_MINUTES;
    }
    public int remainingMinutesInOven(int minutes_in_oven) {
        return expectedMinutesInOven() - minutes_in_oven;
    }
    public int preparationTimeInMinutes(int layers) {
        return layers*LAYER_PREP_TIME;
    }
    public int totalTimeInMinutes(int layers, int minutes) {
        return preparationTimeInMinutes(layers) + minutes;
    }
}
