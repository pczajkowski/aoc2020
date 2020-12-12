#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ROW_SIZE 11
#define ROW_LIMIT 7
#define COLUMN_LIMIT 10
#define ROW_MAX 127
#define COLUMN_MAX 7

int arraySize = 0;

char **readInput(char *path) {
	char **array = NULL;
	FILE *fp = fopen(path, "r");
	if (!fp) {
		puts("Can't read");
		return NULL;
	}

	int index = 0;
	while (1) {
		if (feof(fp)) break;

		char *p = malloc(ROW_SIZE);
		if (!p) return NULL;

		if (1 > fscanf(fp, "%s\n", p))
			return NULL;

		arraySize++;
		char **newArray = realloc(array, sizeof(char*)*arraySize);
		if (!newArray) return NULL;
		array = newArray;
		array[index] = p;
		index++;
	}

	fclose(fp);
	return array;
}

int establishRow(char *code) {
	int rowMin = 0;
	int current = 0;
	int rowMax = ROW_MAX;
	for (int i = 0; i < ROW_LIMIT; i++) {
		if (code[i] == 'F') {
			rowMax = ((rowMax - rowMin) / 2) + rowMin;
			current = rowMin;
		}

		if (code[i] == 'B') {
			rowMin = ((rowMax - rowMin) / 2) + rowMin + 1;
			current = rowMax;
		}
	}

	return current;
}

int establishColumn(char *code) {
	int columnMin = 0;
	int current = 0;
	int columnMax = COLUMN_MAX;
	for (int i = ROW_LIMIT; i < COLUMN_LIMIT; i++) {
		if (code[i] == 'L') {
			columnMax = ((columnMax - columnMin) / 2) + columnMin;
			current = columnMin;
		}

		if (code[i] == 'R') {
			columnMin = ((columnMax - columnMin) / 2) + columnMin + 1;
			current = columnMax;
		}
	}

	return current;
}

int highestID(char **array) {
	int highest = 0;
	for (int i = 0; i < arraySize; i++) {
		int current = (establishRow(array[i]) * 8) + establishColumn(array[i]);
		if (current > highest) highest = current;
	}

	return highest;
}

int **calculateIDs(char **array) {
	int **IDS = malloc(sizeof(int*)*arraySize);
	if (!IDS) return NULL;

	for (int i = 0; i < arraySize; i++) {
		int current = (establishRow(array[i]) * 8) + establishColumn(array[i]);
		IDS[i] = malloc(sizeof(int));
		if (!IDS[i]) return NULL;
		*IDS[i] = current;
	}

	return IDS;
}

int isTaken(int id, int **IDS) {
	for (int i = 0; i < arraySize; i++)
		if (*IDS[i] == id) return 1;

	return 0;
}

int findSeat(int **IDS) {
	for (int i = 0; i < arraySize; i++) {
		for (int j = i + 1; j < arraySize; j++) {
			if ((*IDS[i] - *IDS[j] == 2) || (*IDS[i] - *IDS[j] == -2)) {
				int min = *IDS[i] < *IDS[j] ? *IDS[i] : *IDS[j];
				int seat = min + 1;
				if (!isTaken(seat, IDS)) return seat;
			}
		}

	}
	return 0;
}

void freeArray(char **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

void freeIDS(int **IDS) {
	for (int i = 0; i < arraySize; i++)
		free(IDS[i]);

	free(IDS);
}

int main(int argc, char **argv)
{
	if (argc < 2) {
		puts("You need to specify path to file!");
		return -1;
	}

	char **array = readInput(argv[1]);
	if (!array) {
		puts("Can't read array!");
		return -1;
	}

	printf("Part1: %d\n", highestID(array));

	int **IDS = calculateIDs(array);
	if (!IDS) return -1;

	printf("Part2: %d\n", findSeat(IDS));

	freeArray(array);
	freeIDS(IDS);
}
