const std = @import("std");

pub fn response(s: []const u8) []const u8 {
    const trimmed = std.mem.trim(u8, s, "\t\r\n ");
    const isQuestion = std.mem.endsWith(u8, trimmed, "?");
    const isSilence = std.mem.eql(u8, trimmed, "");
    const isAllCap = hasAllCap(trimmed);

    if (isQuestion and isAllCap) {
        return "Calm down, I know what I'm doing!";
    } else if (isQuestion) {
        return "Sure.";
    } else if (isAllCap) {
        return "Whoa, chill out!";
    } else if (isSilence) {
        return "Fine. Be that way!";
    } else {
        return "Whatever.";
    }
}

fn hasAllCap(s: []const u8) bool {
    var hasAlpha = false;
    for (s) |c| {
        if (std.ascii.isAlphabetic(c)) {
            hasAlpha = true;
            if (std.ascii.isLower(c)) {
                return false;
            }
        }
    }
    return hasAlpha;
}
