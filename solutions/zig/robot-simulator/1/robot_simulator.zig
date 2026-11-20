pub const Direction = enum {
    north,
    east,
    south,
    west,
};

pub const Robot = struct {
    // This struct, as well as its fields and methods, needs to be implemented.
    x: i32,
    y: i32,
    direction: Direction,

    pub fn init(x: i32, y: i32, direction: Direction) Robot {
        return Robot{ .x = x, .y = y, .direction = direction };
    }

    pub fn move(self: *Robot, instructions: []const u8) void {
        for (instructions) |ins| {
            switch (ins) {
                'R' => self.right(),
                'L' => self.left(),
                'A' => self.advance(),
                else => unreachable,
            }
        }
    }

    fn right(self: *Robot) void {
        switch (self.direction) {
            Direction.north => self.direction = Direction.east,
            Direction.east => self.direction = Direction.south,
            Direction.south => self.direction = Direction.west,
            Direction.west => self.direction = Direction.north,
        }
    }

    fn left(self: *Robot) void {
        switch (self.direction) {
            Direction.north => self.direction = Direction.west,
            Direction.east => self.direction = Direction.north,
            Direction.south => self.direction = Direction.east,
            Direction.west => self.direction = Direction.south,
        }
    }

    fn advance(self: *Robot) void {
        switch (self.direction) {
            Direction.north => self.y += 1,
            Direction.east => self.x += 1,
            Direction.south => self.y -= 1,
            Direction.west => self.x -= 1,
        }
    }
};
