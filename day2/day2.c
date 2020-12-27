#include <stdio.h>

#define PASS_MAX 50

void part1(FILE *fp) {
	int validCount = 0;

	while (1) {
		if (feof(fp)) break;

		int from, to;
		char check;
		char password[PASS_MAX];

		if (4 != fscanf(fp, "%d-%d %c: %s\n", &from, &to, &check, password))
			return;

		int index = 0;
		int checkCount = 0;
		while (1) {
			if (password[index] == '\0') break;

			if (password[index] == check) checkCount++;

			if (checkCount > to) break;

			index++;
		}

		if (checkCount >= from && checkCount <= to)
			validCount++;
	}

	printf("Part1: %d\n", validCount);
}

void part2(FILE *fp) {
	int validCount = 0;

	while (1) {
		if (feof(fp)) break;

		int from, to;
		char check;
		char password[PASS_MAX];

		if (4 != fscanf(fp, "%d-%d %c: %s\n", &from, &to, &check, password))
			return;

		if (password[to-1] == check ^ password[from-1] == check) validCount++;
	}

	printf("Part2: %d\n", validCount);
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

	part1(fp);
	rewind(fp);
	part2(fp);
	fclose(fp);
}
