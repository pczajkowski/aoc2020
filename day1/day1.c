#include <stdio.h>
#include <stdlib.h>

int arraySize = 0;

int **readInput(FILE *fp) {
	int **array = NULL;
	int index = 0;

	while (1) {
		if (feof(fp)) break;

		int *p = malloc(sizeof(int));
		if (!p) return NULL;

		if (1 != fscanf(fp, "%d\n", p))
			return NULL;

		arraySize++;
		int **newArray = realloc(array, sizeof(int*)*arraySize);
		if (!newArray) return NULL;
		array = newArray;
		array[index] = p;
		index++;
	}

	return array;
}

void freeArray(int **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

void part1(int **array, int target) {
	for (int i = 0; i < arraySize; i++) {
		int first = *array[i];
		int second = target - first;

		for (int j = i + 1; j < arraySize; j++) {
			if (*array[j] == second) {
				printf("Part1: %d\n", first * second);
				return;
			}
		}
	}
}

void part2(int **array, int target) {
	for (int i = 0; i < arraySize; i++) {
		int first = *array[i];
		int threshold = target - first;

		for (int j = i + 1; j < arraySize; j++) {
			if (*array[j] < threshold) {
				int third = threshold - *array[j];
				for (int k = j + 1; k < arraySize; k++) {
					if (*array[k] == third) {
						printf("Part1: %d\n", first * *array[j] *third);
						return;
					}

				}
			}
		}
	}
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

	int **array = readInput(fp);
	if (!array) {
		puts("Can't read array!");
		return -1;
	}

	part1(array, 2020);
	part2(array, 2020);
	fclose(fp);
	freeArray(array);
}
