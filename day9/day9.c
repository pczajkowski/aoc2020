#include <stdio.h>
#include <stdlib.h>

int arraySize = 0;

int checkNumber(int number, int **array, int last) {
	for (int i = 0; i < last; i++)
		for (int j = i + 1; j < last; j++) {
			if ((*array[i] + *array[j]) == number) {
				return 1;
			}
		}

	return 0;
}

int findNumber(int **allNumbers, int last) {
	int tempArray[last];
	for (int i = 0; i < last; i++)
		tempArray[i] = *allNumbers[i];

	int start = 0;
	int index = last;
	while (index < arraySize) {
		int number = *allNumbers[index];
		if (!checkNumber(number, &allNumbers[start++], last)) return number;
		index++;
	}

	return -1;
}

int **readAllNumbers(FILE *fp) {
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

int encryptionWeakness(int badNumber, int **allNumbers) {
	for (int i = 0; i < arraySize; i++) {
		int sum = *allNumbers[i];
		int smallest = *allNumbers[i];
		int largest = *allNumbers[i];
		for (int j = i + 1; j < arraySize; j++) {
			sum += *allNumbers[j];
			smallest = smallest < *allNumbers[j] ? smallest : *allNumbers[j];
			largest = largest > *allNumbers[j] ? largest : *allNumbers[j];
			if (sum == badNumber) return smallest + largest;
		}
	}

	return -1;
}

void freeAllNumbers(int **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

int main(int argc, char **argv)
{
	if (argc < 3) {
		puts("You need to specify path to file and preamble length!");
		return -1;
	}

	int last = atoi(argv[2]);
	FILE *fp = fopen(argv[1], "r");
	if (!fp) {
		puts("Can't read");
		return -1;
	}

	int **allNumbers = readAllNumbers(fp);
	if (!allNumbers) {
		puts("Can't read numbers!");
		fclose(fp);
		return -1;
	}

	int badNumber = findNumber(allNumbers, last);
	printf("Part1: %d\n", badNumber);

	printf("Part2: %d\n", encryptionWeakness(badNumber, allNumbers));

	freeAllNumbers(allNumbers);
	fclose(fp);
}
