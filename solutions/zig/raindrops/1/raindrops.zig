const std = @import("std");

pub fn convert(buffer: []u8, n: u32) []const u8 {
    var len: usize = 0;

    if (n % 3 == 0) {
        @memcpy(buffer[len .. len + 5], "Pling");
        len += 5;
    }
    if (n % 5 == 0) {
        @memcpy(buffer[len .. len + 5], "Plang");
        len += 5;
    }
    if (n % 7 == 0) {
        @memcpy(buffer[len .. len + 5], "Plong");
        len += 5;
    }

    if (len == 0) {
        return std.fmt.bufPrint(buffer, "{d}", .{n}) catch unreachable;
    }
    return buffer[0..len];
}
