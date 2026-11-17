const std = @import("std");

pub fn truncate(phrase: []const u8) []const u8 {
    const view = std.unicode.Utf8View.init(phrase) catch return phrase;
    var it = view.iterator();
    for (0..5) |_| {
        _ = it.nextCodepoint() orelse return phrase;
    }
    return phrase[0..it.i];
}
