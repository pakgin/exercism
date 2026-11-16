const std = @import("std");
const mem = std.mem;

pub const Node = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    data: i32,
    left: ?*Node = null,
    right: ?*Node = null,
};

pub const Tree = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    allocator: mem.Allocator,
    root: ?*Node,

    pub fn init(allocator: mem.Allocator) Tree {
        return .{ .allocator = allocator, .root = null };
    }

    pub fn deinit(self: *Tree) void {
        destroy(self.allocator, self.root);
        self.root = null;
    }

    fn destroy(allocator: mem.Allocator, root: ?*Node) void {
        const n = root orelse return;
        destroy(allocator, n.left);
        destroy(allocator, n.right);
        allocator.destroy(n);
    }

    pub fn insert(self: *Tree, data: i32) mem.Allocator.Error!void {
        const new = try self.allocator.create(Node);
        new.* = .{ .data = data };

        var cursor = &self.root;
        while (cursor.*) |n| {
            cursor = if (data <= n.data) &n.left else &n.right;
        }
        cursor.* = new;
    }

    pub fn sortedData(self: *const Tree, allocator: mem.Allocator) mem.Allocator.Error![]i32 {
        var list: std.ArrayList(i32) = .empty;
        errdefer list.deinit(allocator);

        try collect(allocator, &list, self.root);
        return list.toOwnedSlice(allocator);
    }

    fn collect(allocator: mem.Allocator, list: *std.ArrayList(i32), root: ?*Node) mem.Allocator.Error!void {
        const n = root orelse return;
        try collect(allocator, list, n.left);
        try list.append(allocator, n.data);
        try collect(allocator, list, n.right);
    }
};
