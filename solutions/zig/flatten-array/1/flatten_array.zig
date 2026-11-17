const std = @import("std");
const mem = std.mem;

pub const Box = union(enum) {
    none,
    one: i12,
    many: []const Box,
};

pub fn flatten(allocator: mem.Allocator, box: Box) mem.Allocator.Error![]i12 {
    // _ = allocator;
    // _ = box;
    // @compileError("please implement the flatten function");

    var list: std.ArrayList(i12) = .empty;
    defer list.deinit(allocator);

    try _flatten(&list, allocator, box);
    return list.toOwnedSlice(allocator);
}

fn _flatten(list: *std.ArrayList(i12), allocator: mem.Allocator, box: Box) mem.Allocator.Error!void {
    switch (box) {
        .none => {},
        .one => |n| try list.append(allocator, n),
        .many => |boxes| {
            for (boxes) |inner| {
                try _flatten(list, allocator, inner);
            }
        },
    }
}
