#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NO_ANSWERS 26

int countChecks1(int *checks) {
	int count = 0;
	for (int i = 0; i < NO_ANSWERS; i++) {
		if (checks[i] > 0) count++;
	}

	return count;
}

int countChecks2(int *checks, int groupSize) {
	int count = 0;
	for (int i = 0; i < NO_ANSWERS; i++) {
		if (checks[i] == groupSize) count++;
	}

	return count;
}

void zeroChecks(int *checks) {
	for (int i = 0; i < NO_ANSWERS; i++)
		checks[i] = 0;
}

int readInput(char *path) {
	FILE *fp = fopen(path, "r");
	if (!fp) {
		puts("Can't read");
		return -1;
	}

	int checks[NO_ANSWERS];
	zeroChecks(checks);

	int sum1 = 0;
	int sum2 = 0;
	int groupSize = 0;
	while (1) {
		char * line = NULL;
		size_t len = 0;
		ssize_t length = getline(&line, &len, fp);

		if ((length == 1 && line[0] == '\n') || feof(fp)) {
			sum1 += countChecks1(checks);
			sum2 += countChecks2(checks, groupSize);
			zeroChecks(checks);
			groupSize = 0;
		} else {
			groupSize++;
			for (int i = 0; i < length-1; i++) {
				int index = line[i]-97;
				checks[index]++;
			}
		}

		if (feof(fp)) break;
	}

	printf("Part1: %d\n", sum1);
	printf("Part2: %d\n", sum2);
	return 1;
}

int main(int argc, char **argv)
{
	if (argc < 2) {
		puts("You need to specify path to file!");
		return -1;
	}

	if (!readInput(argv[1])) {
		puts("Can't read array!");
		return -1;
	}
}
