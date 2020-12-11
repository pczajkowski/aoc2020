#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int arraySize = 0;

struct instruction {
	char name[4];
	int value;
};

void zeroArray(int *array) {
	for (int i = 0; i < arraySize; i++)
		array[i] = 0;
}

struct instruction **readInput(char *path) {
	struct instruction **instructions = NULL;
	FILE *fp = fopen(path, "r");
	if (!fp) {
		puts("Can't read");
		return NULL;
	}

	int index = 0;
	while (1) {
		struct instruction *p = malloc(sizeof(struct instruction));
		if (!p) return NULL;

		if (1 > fscanf(fp, "%s %d\n", p->name, &(p->value)))
			return NULL;


		arraySize++;
		struct instruction **newInstructions = realloc(instructions, sizeof(struct instruction*)*arraySize);
		if (!newInstructions) return NULL;
		instructions = newInstructions;
		instructions[index] = p;
		index++;

		if (feof(fp)) break;
	}

	fclose(fp);
	return instructions;
}

int processInstructions(struct instruction **instructions, int *executed) {
	int acc = 0;
	int index = 0;
	while(1) {
		if (index < 0 || index >= arraySize) break;
		if (executed[index]) break;

		struct instruction *currentInstruction = instructions[index];
		if (0 == (strcmp("nop", currentInstruction->name))) {
			executed[index] = 1;
			index++;
			continue;
		}

		if (0 == (strcmp("acc", currentInstruction->name))) {
			acc += currentInstruction->value;
			executed[index] = 1;
			index++;
			continue;
		}

		if (0 == (strcmp("jmp", currentInstruction->name))) {
			executed[index] = 1;
			index += currentInstruction->value;
			continue;
		}
	}

	return acc;
}

int checkArray(struct instruction **instructions, int *executed, int index) {
	int executedCheck[arraySize];
	zeroArray(executedCheck);

	while(1) {
		if (index < 0 || index >= arraySize) break;
		if (executed[index] || executedCheck[index]) break;

		struct instruction *currentInstruction = instructions[index];
		if (0 == (strcmp("nop", currentInstruction->name))) {
			executedCheck[index] = 1;
			index++;
			continue;
		}

		if (0 == (strcmp("acc", currentInstruction->name))) {
			executedCheck[index] = 1;
			index++;
			continue;
		}

		if (0 == (strcmp("jmp", currentInstruction->name))) {
			executedCheck[index] = 1;
			index += currentInstruction->value;
			continue;
		}
	}

	return index == arraySize;
}

int fixTheLoop(struct instruction **instructions, int *executed) {
	int acc = 0;
	int index = 0;
	int changed = 0;
	while(1) {
		if (index < 0 || index >= arraySize) break;
		if (executed[index]) break;

		struct instruction *currentInstruction = instructions[index];
		if (0 == (strcmp("nop", currentInstruction->name))) {
			executed[index] = 1;

			if (!changed && checkArray(instructions, executed, index+currentInstruction->value)) {
				changed = 1;
				index += currentInstruction->value;
			}
			else index++;

			continue;
		}

		if (0 == (strcmp("acc", currentInstruction->name))) {
			acc += currentInstruction->value;
			executed[index] = 1;
			index++;
			continue;
		}

		if (0 == (strcmp("jmp", currentInstruction->name))) {
			executed[index] = 1;

			if (!changed && checkArray(instructions, executed, index+1)) {
				changed = 1;
				index++;
			}

			else index += currentInstruction->value;

			continue;
		}
	}

	if (changed) return acc;
	return 1000000000;
}

void freeInstructions(struct instruction **instructions) {
	for (int i = 0; i < arraySize; i++)
		free(instructions[i]);

	free(instructions);
}

int main(int argc, char **argv)
{
	if (argc < 2) {
		puts("You need to specify path to file!");
		return -1;
	}

	struct instruction **instructions = readInput(argv[1]);
	if (!instructions) {
		puts("Can't read array!");
		return -1;
	}

	int executed[arraySize];
	zeroArray(executed);

	printf("Part1: %d\n", processInstructions(instructions, executed));

	zeroArray(executed);
	printf("Part2: %d\n", fixTheLoop(instructions, executed));

	freeInstructions(instructions);
}
