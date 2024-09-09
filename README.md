# Go Algorithm katas

## Daily Kata Practice

To facilitate daily algorithm practice, we've created a script called `practice_kata.sh`. This script allows you to generate new practice days and copy specific katas to practice.

### Usage

1. Generate a new practice day:

```bash
./practice_kata.sh new-day
```

This creates a new directory in the `practice` folder with the name `day_X`, where X is the next available number.

2. Copy a kata to the latest practice day:

```bash
./practice_kata.sh copy-kata binary_tree/avl_tree
```

This copies the AVL tree kata template to the latest practice day.

3. Copy a kata to a specific practice day:

```bash
./practice_kata.sh copy-kata graph/dijkstra day_3
```

This copies the Dijkstra's algorithm kata template to the `day_3` directory.

### Practice Structure

After using this script, your practice structure will look something like this:

```
practice/
├── day_1/
│   ├── avl_tree/
│   │   ├── avl_tree.go
│   │   ├── avl_tree_test.go
│   │   └── README.md
│   └── ...
├── day_2/
│   ├── dijkstra/
│   │   ├── dijkstra.go
│   │   ├── dijkstra_test.go
│   │   └── README.md
│   └── ...
└── ...
```

### Notes

- The script assumes that your kata templates are in the `templates` directory.
- Run the script from the root of your project.
- The script prevents overwriting existing katas in a practice day to avoid accidental data loss.

By using this script, you can easily set up daily practice sessions, allowing you to implement the same algorithm multiple times to improve your skills and track your progress over time.
