#include <stdio.h>
#include <stdlib.h>

int arraySize = 0;

int **readAllJolts(FILE *fp) {
	int **allNumbers = NULL;
	int index = 0;
	while (1) {
		int *p = malloc(sizeof(int));
		if (!p) return NULL;

		if (1 > fscanf(fp, "%d\n", p))
			return NULL;

		arraySize++;
		int **newNumbers = realloc(allNumbers, sizeof(int*)*arraySize);
		if (!newNumbers) return NULL;

		allNumbers = newNumbers;
		allNumbers[index] = p;
		index++;

		if (feof(fp)) break;
	}

	return allNumbers;
}

int compare(const void *el1, const void *el2) {
	int *value1 = *(int**)el1;
	int *value2 = *(int**)el2;
	return (*value1 < *value2) ? -1 : (*value1 > *value2) ? 1 : 0;
}

char *calculateDistribution(int **array) {
	char *sequence = malloc(arraySize+1);
	if (!sequence) return NULL;
	int sequenceIndex = 0;

	int start = 0;
	int oneJolt = 0;
	int threeJolt = 0;

	int smallestIndex = 0;
	for (int i = 0; i < arraySize; i++) {
		if (*array[i] <= (start+3)) {
			int difference = i > 0 ? *array[i] - *array[i-1] : *array[i] - start;

			switch (difference) {
				case 1:
					sequence[sequenceIndex] = '1';
					oneJolt++;
					break;
				case 2:
					sequence[sequenceIndex] = '1';
					sequenceIndex++;
					sequence[sequenceIndex] = '1';
					break;
				case 3:
					sequence[sequenceIndex] = 'd';
					threeJolt++;
					break;
			}

			sequenceIndex++;

			if (smallestIndex != i) {
				if (*array[smallestIndex] > *array[i]) {
					smallestIndex = i;
				}
			}
		} else {
			start = *array[smallestIndex];


			smallestIndex = i;
			i--;
		}
	}

	threeJolt++;
	printf("Part1: %d\n", oneJolt * threeJolt);

	sequence[arraySize] = '\0';
	return sequence;
}

void freeAllNumbers(int **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

long int processSequence(char *sequence) {
	long int result = 1;

	int ones = 0;
	for (int i = 0; i < arraySize+1; i++) {
		if (sequence[i] == '1') ones++;

		if (sequence[i] == 'd' || sequence[i] == '\0') {
			switch (ones) {
				case 2:
					result *= 2;
					break;
				case 3:
					result *= 4;
					break;
				case 4:
					result *= 7;
					break;
			}
			ones = 0;
		}
	}

	return result;
}

int main(int argc, char **argv)
{
	if (argc < 2) {
		puts("You need to specify path to file!");
		return -1;
	}

	FILE *fp = fopen(argv[1], "r");
	if (!fp) {
		puts("Can't open file!");
		return -1;
	}

	int **allJolts = readAllJolts(fp);
	if (!allJolts) {
		puts("Can't read numbers!");
		fclose(fp);
		return -1;
	}

	qsort(allJolts, arraySize, sizeof(int*), compare);
	char *sequence = calculateDistribution(allJolts);
	printf("Part2: %ld\n", processSequence(sequence));

	free(sequence);
	freeAllNumbers(allJolts);
	fclose(fp);
}
