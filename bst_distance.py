class Node:
    def __init__(self, val):
        self.value = val
        self.right = None
        self.left = None
        self.visited = False
    
    def __str__(self):
        r = None if self.right is None else self.right.value
        l = None if self.left is None else self.left.value
        return f"({self.value}): left: {l} | right: {r}"

class Tree:
    def __init__(self):
        self.head = None

    def reset_visited(self):
        if self.head is None:
            return
        queue = [self.head]
        while queue:
            n = queue.pop()
            n.visited = False
            if n.left is not None:
                queue.append(n.left)
            if n.right is not None:
                queue.append(n.right)
    
    def print(self):
        if self.head is None:
            return
        queue = [self.head]
        lvl = 0
        while queue:
            lvl += 1
            print(f"Level: {lvl}")
            items = len(queue)
            for _ in range(items):
                n = queue.pop(0)
                print(n)
                if n.left is not None:
                    queue.append(n.left)
                if n.right is not None:
                    queue.append(n.right)
    
    def parse_bst(self, arr):
        self.head = None
        self.head = Node(arr[0])
        for elem in arr[1:]:
            n = self.head
            placed = False
            while n is not None:
                side = 'left' if elem < n.value else 'right'
                child = getattr(n, side)
                if child is None:
                    setattr(n, side, Node(elem))
                n = child
    
    def distance_between_nodes_bst(self, a, b):
        def distance_to(ancestor, num):
            dist = 0
            curr = ancestor
            while curr is not None and curr.value != num:
                dist += 1
                curr = curr.left if num < curr.value else curr.right
            if curr is None:
                return -1
            return dist
        if a == b:
            return 0
        ancestor = None
        queue = [self.head]
        while ancestor is None:
            nxt = queue.pop(0)
            v = nxt.value
            if v == a or v == b or a < v < b or b < v < a:
                ancestor = nxt
                queue = []
        dist_a = distance_to(ancestor, a)
        dist_b = distance_to(ancestor, b)
        if dist_a == -1 or dist_b == -1:
            return -1
        return dist_a + dist_b
        

    def distance_between_nodes_dfs(self, a, b):
        if self.head is None:
            return
        stack = [self.head]
        found_a = False
        found_b = False
        path_a = [self.head.value]
        path_b = [self.head.value]
        while stack and (not found_a or not found_b):
            n = stack.pop()
            print(n)
            if not found_a:
                path_a.pop()
            if not found_b:
                path_b.pop()
            while n is not None and not n.visited:
                n.visited = True
                stack.append(n)
                if not found_a:
                    path_a.append(n.value)
                    found_a = path_a[-1] == a
                if not found_b:
                    path_b.append(n.value)
                    found_b = path_b[-1] == b
                n = n.left
            n = stack[-1]
            if n.right is not None and not n.right.visited:
                stack.append(n.right)
                if not found_a:
                    path_a.append(None)
                if not found_b:
                    path_b.append(None)
        print(path_a)
        print(path_b)
        min_size = min(len(path_b), len(path_a))
        i = 0
        while i < min_size and path_a[i] == path_b[i]:
            i += 1
        return (len(path_a) - i) + (len(path_b) - i)


                


t = Tree()
t.parse_bst([5,7,3,6,2,1,4,8])

print(t.distance_between_nodes_bst(5, 8))
print(t.distance_between_nodes_bst(2, 8))
print(t.distance_between_nodes_bst(1, 6))