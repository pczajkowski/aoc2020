#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINE_MAX 50

int arraySize = 0;

char **readInput(FILE *fp) {
	char **array = NULL;
	int index = 0;

	while (1) {
		if (feof(fp)) break;

		char *p = malloc(LINE_MAX);
		if (!p) return NULL;

		if (1 != fscanf(fp, "%s\n", p))
			return NULL;

		arraySize++;
		char **newArray = realloc(array, sizeof(char*)*arraySize);
		if (!newArray) return NULL;
		array = newArray;
		array[index] = p;
		index++;
	}

	return array;
}

void freeArray(char **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

long int traverse(int right, int down, char **array) {
	long int trees = 0;

	size_t x = right;
	size_t xMax = strlen(array[0]) - 1;
	int y = down;

	while (y < arraySize) {
		if (x > xMax) x = x - xMax - 1;

		if (array[y][x] == '#') trees++;

		x += right;
		y += down;
	}

	return trees;
}

int main(int argc, char **argv)
{
	if (argc < 2) {
		puts("You need to specify path to file!");
		return -1;
	}

	FILE *fp = fopen(argv[1], "r");
	if (!fp) {
		puts("Can't read");
		return -1;
	}

	char **array = readInput(fp);
	if (!array) {
		puts("Can't read array!");
		return -1;
	}

	long int part1 = traverse(3, 1, array);
	printf("Part1: %ld\n", part1);

	long int x1 = traverse(1, 1, array);
	long int x2 = traverse(5, 1, array);
	long int x3 = traverse(7, 1, array);
	long int x4 = traverse(1, 2, array);

	printf("Part2: %ld\n", part1*x1*x2*x3*x4);

	fclose(fp);
	freeArray(array);
}
