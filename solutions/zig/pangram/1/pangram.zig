const std = @import("std");

pub fn isPangram(str: []const u8) bool {
    var counter = std.mem.zeroes([26]u8);
    for (str) |ch| {
        if (!std.ascii.isAlphabetic(ch)) continue;
        counter[std.ascii.toLower(ch) - 'a'] += 1;
    }

    for (counter) |c| {
        if (c == 0) {
            return false;
        }
    }
    return true;
}
