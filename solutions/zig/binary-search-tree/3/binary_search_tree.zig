const std = @import("std");
const mem = std.mem;

pub const Node = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    data: i32,
    left: ?*Node,
    right: ?*Node,

    pub fn init(allocator: mem.Allocator, data: i32) mem.Allocator.Error!*Node {
        const self = try allocator.create(Node);
        self.* = .{ .data = data, .left = null, .right = null };
        return self;
    }

    pub fn deinit(self: *Node, allocator: mem.Allocator) void {
        if (self.left) |l| l.deinit(allocator);
        if (self.right) |r| r.deinit(allocator);
        allocator.destroy(self);
    }

    pub fn insert(self: *Node, allocator: mem.Allocator, data: i32) mem.Allocator.Error!void {
        if (data <= self.data) {
            if (self.left) |child| {
                try child.insert(allocator, data);
            } else {
                self.left = try Node.init(allocator, data);
            }
        } else {
            if (self.right) |child| {
                try child.insert(allocator, data);
            } else {
                self.right = try Node.init(allocator, data);
            }
        }
    }
};

pub const Tree = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    allocator: mem.Allocator,
    root: ?*Node,

    pub fn init(allocator: mem.Allocator) Tree {
        return .{ .allocator = allocator, .root = null };
    }

    pub fn deinit(self: *Tree) void {
        const n = self.root orelse return;
        n.deinit(self.allocator);
        self.root = null;
    }

    pub fn insert(self: *Tree, data: i32) mem.Allocator.Error!void {
        if (self.root) |child| {
            try child.insert(self.allocator, data);
        } else {
            self.root = try Node.init(self.allocator, data);
        }
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
