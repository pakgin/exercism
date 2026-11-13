import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class LogLevels {
    private static final String REGEX = "\\[(?<level>\\w+)\\]: (?<message>[\\w\\s]+)";
    private static final Pattern PATTERN = Pattern.compile(REGEX);

    private static Matcher match(String logLine) {
        var matcher = PATTERN.matcher(logLine);
        matcher.find();
        return matcher;
    }
    
    public static String message(String logLine) {
        return match(logLine).group("message").trim();
    }

    public static String logLevel(String logLine) {
        return match(logLine).group("level").toLowerCase();
    }

    public static String reformat(String logLine) {
        return String.format("%s (%s)", message(logLine), logLevel(logLine));
    }
}
