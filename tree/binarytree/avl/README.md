# 平衡二叉树


### LL, right rotation:
```
        x                   y
       /                   / \
      y        -->        z   x
     /
    z
 ```

### RR, left rotation:
```
    x                       y
     \                     / \
      y        -->        x   z
       \
        z
```

### LR, left right rotation:
```
        x                   x                     z
       /                   /                     / \
      y        -->        z         -->         y   x
       \                 /
        z               y
```

### RL, right left rotation:
```
    x                   x                         z
     \                   \                       / \
      y        -->        z         -->         x   y
     /                     \
    z                       y
```