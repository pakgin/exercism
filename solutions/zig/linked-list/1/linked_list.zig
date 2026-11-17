const std = @import("std");

pub fn LinkedList(comptime T: type) type {
    return struct {
        const Self = @This();

        // Please implement the doubly linked `Node` (replacing each `void`).
        pub const Node = struct {
            prev: ?*Node = null,
            next: ?*Node = null,
            data: T,
        };

        // Please implement the fields of the linked list (replacing each `void`).
        first: ?*Node = null,
        last: ?*Node = null,
        len: usize = 0,

        // Please implement the below methods.
        // You need to add the parameters to each method.

        /// Add a node at the back.
        pub fn push(self: *Self, elem: *Node) void {
            elem.prev = self.last;
            elem.next = null;
            if (self.last) |last| {
                last.next = elem;
            } else {
                self.first = elem; // list was empty
            }
            self.last = elem;
            self.len += 1;
        }

        /// Remove and return the back node, or null if empty.
        pub fn pop(self: *Self) ?*Node {
            const n = self.last orelse return null;
            self.last = n.prev;
            if (self.last) |l| {
                l.next = null;
            } else {
                self.first = null; // list is now empty
            }
            n.prev = null;
            n.next = null;
            self.len -= 1;
            return n;
        }

        /// Remove and return the front node, or null if empty.
        pub fn shift(self: *Self) ?*Node {
            const n = self.first orelse return null;
            self.first = n.next;
            if (self.first) |f| {
                f.prev = null;
            } else {
                self.last = null; // list is now empty
            }
            n.prev = null;
            n.next = null;
            self.len -= 1;
            return n;
        }

        pub fn unshift(self: *Self, elem: *Node) void {
            elem.next = self.first;
            elem.prev = null;
            if (self.first) |first| {
                first.prev = elem;
            } else {
                self.last = elem; // list was empty
            }
            self.first = elem;
            self.len += 1;
        }

        pub fn delete(self: *Self, elem: *Node) void {
            var cursor = self.first;
            while (cursor) |c| : (cursor = c.next) {
                if (c != elem) continue;

                if (c.prev) |p| {
                    p.next = c.next;
                } else {
                    self.first = c.next; // removing the head
                }
                if (c.next) |n| {
                    n.prev = c.prev;
                } else {
                    self.last = c.prev; // removing the tail
                }
                c.prev = null;
                c.next = null;
                self.len -= 1;
                return;
            }
        }
    };
}
