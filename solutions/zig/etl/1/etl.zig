const std = @import("std");
const mem = std.mem;

pub fn transform(allocator: mem.Allocator, legacy: std.AutoHashMap(i5, []const u8)) mem.Allocator.Error!std.AutoHashMap(u8, i5) {
    var new = std.AutoHashMap(u8, i5).init(allocator);
    errdefer new.deinit();

    var iter = legacy.iterator();
    while (iter.next()) |entry| {
        const key, const values = .{ entry.key_ptr, entry.value_ptr };
        for (values.*) |v| {
            try new.put(std.ascii.toLower(v), key.*);
        }
    }
    return new;
}
