import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class LogLevels {
    static final Pattern PATTERN;
    static {
        String pattern = "\\[(\\w+)\\]: ([\\w\\s]+)";
        PATTERN = Pattern.compile(pattern);
    }
    
    public static String message(String logLine) {
        Matcher m = PATTERN.matcher(logLine);
        if (m.find()) {
            return m.group(2).trim();
        }
        return "unknown";
    }

    public static String logLevel(String logLine) {
        Matcher m = PATTERN.matcher(logLine);
        if (m.find()) {
            return m.group(1).toLowerCase();
        }
        return "error";
    }

    public static String reformat(String logLine) {
        Matcher m = PATTERN.matcher(logLine);
        if (m.find()) {
            return String.format("%s (%s)", message(logLine), logLevel(logLine));
        }
        return "error";
    }
}
