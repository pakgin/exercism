const std = @import("std");
const mem = std.mem;

pub const TraversalError = error{
    DifferentLengths,
    DifferentItems,
    NonUniqueItems,
};

pub const Node = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    data: u8,
    left: ?*Node,
    right: ?*Node,

    pub fn init(allocator: mem.Allocator, data: u8, left: ?*Node, right: ?*Node) mem.Allocator.Error!*Node {
        const self = try allocator.create(Node);
        errdefer allocator.destroy(self);
        self.* = .{ .data = data, .left = left, .right = right };
        return self;
    }

    pub fn deinit(self: *Node, allocator: mem.Allocator) void {
        if (self.left) |left| {
            left.deinit(allocator);
            allocator.destroy(left);
        }
        if (self.right) |right| {
            right.deinit(allocator);
            allocator.destroy(right);
        }
        self.left = null;
        self.right = null;
    }
};

pub const Tree = struct {
    // This struct, as well as its fields and methods, needs to be implemented.

    allocator: mem.Allocator,
    root: ?*Node,

    pub fn initFromTraversals(allocator: mem.Allocator, preorder: []const u8, inorder: []const u8) (mem.Allocator.Error || TraversalError)!Tree {
        if (preorder.len != inorder.len) return TraversalError.DifferentLengths;

        var seen = std.AutoHashMap(u8, void).init(allocator);
        defer seen.deinit();

        return .{ .allocator = allocator, .root = try collect(allocator, &seen, preorder, inorder) };
    }

    fn collect(allocator: mem.Allocator, seen: *std.AutoHashMap(u8, void), preorder: []const u8, inorder: []const u8) (mem.Allocator.Error || TraversalError)!?*Node {
        if (preorder.len == 0 or inorder.len == 0) return null;

        const root = preorder[0];
        if (seen.contains(root)) {
            return TraversalError.NonUniqueItems;
        }
        try seen.put(root, {});

        const index = std.mem.indexOfScalar(u8, inorder, root) orelse return TraversalError.DifferentItems;
        const left = try collect(allocator, seen, preorder[1 .. index + 1], inorder[0..index]);
        errdefer if (left) |l| {
            l.deinit(allocator);
            allocator.destroy(l);
        };

        const right = try collect(allocator, seen, preorder[1 + index ..], inorder[index + 1 ..]);
        errdefer if (right) |r| {
            r.deinit(allocator);
            allocator.destroy(r);
        };

        return try Node.init(allocator, root, left, right);
    }

    pub fn deinit(self: *Tree) void {
        if (self.root) |root| {
            root.deinit(self.allocator);
            self.allocator.destroy(root);
        }
        self.root = null;
    }
};
