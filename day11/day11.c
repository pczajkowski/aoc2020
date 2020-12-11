#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LINE_MAX 256

int arraySize = 0;

char **readSeats(FILE *fp) {
	char **allSeats = NULL;
	int index = 0;
	while (1) {
		char *p = malloc(LINE_MAX);
		if (!p) return NULL;

		if (1 > fscanf(fp, "%s\n", p))
			return NULL;

		arraySize++;
		char **newSeats = realloc(allSeats, sizeof(char*)*arraySize);
		if (!newSeats) return NULL;

		allSeats = newSeats;
		allSeats[index] = p;
		index++;

		if (feof(fp)) break;
	}

	return allSeats;
}

size_t lineLength = 0;

int occupied(char **array, int x, int y) {
	int columnMax = lineLength-1;
	int rowMax = arraySize-1;
	if ((x < 0 || x > rowMax) || (y < 0 || y > columnMax)) return 0;

	int upRow = x - 1;
	if (upRow < 0) upRow = 0;
	int downRow = x + 1;
	if (downRow > rowMax) downRow = rowMax;

	int left = y - 1;
	if (left < 0) left = 0;
	int right = y + 1;
	if (right > columnMax) right = columnMax;

	int occupiedCount = 0;
	for (int i = upRow; i <= downRow; i++) {
		for (int j = left; j <= right; j++) {
			if (i == x && j == y) continue;
			if (array[i][j] == '#') occupiedCount++;
		}
	}

	return occupiedCount;
}



void freeAllSeats(char **array) {
	for (int i = 0; i < arraySize; i++)
		free(array[i]);

	free(array);
}

int previouslyChanged = 0;
int currentlyChanged = 0;

char **checkBoard(char **currentBoard, char **newBoard) {
	currentlyChanged = 0;
	for (int i = 0; i < arraySize; i++) {
		char temp[lineLength+1];
		temp[lineLength] = 0;
		for (size_t j = 0; j < lineLength; j++) {
			char *current = &currentBoard[i][j];
			char new = *current;

			switch (*current) {
				case 'L':
					if (!occupied(currentBoard, i, j)) {
						new = '#';
						currentlyChanged = 1;
					}
					break;
				case '#':
					if (occupied(currentBoard, i, j) >= 4) {
						new = 'L';
						currentlyChanged = 1;
					}
					break;
			}

			temp[j] = new;
		}
		sprintf(newBoard[i], "%s", temp);
	}

	freeAllSeats(currentBoard);

	return newBoard;
}

int countOccupied(char **currentBoard) {
	int occupiedCount = 0;
	for (int i = 0; i < arraySize; i++) {
		for (size_t j = 0; j < lineLength; j++) {
			char *current = &currentBoard[i][j];
			if (*current == '#') {
				occupiedCount++;
			}
		}
	}

	return occupiedCount;
}

char **allocateNewBoard(int rows, int columns) {
	char **newBoard = malloc(sizeof(char*)*rows);
	if (!newBoard) return NULL;

	for (int i = 0; i < rows; i++) {
		newBoard[i] = malloc(columns);
		if (!newBoard[i]) return NULL;
	}

	return newBoard;
}


int doRounds(char **array) {
	char **currentBoard = array;
	int runs = 0;
	char **newBoard = NULL;
	while (1) {
		newBoard = allocateNewBoard(arraySize, lineLength+1);
		if (!newBoard) return -1;

		if (runs > 1 && ((currentlyChanged == 0) && (previouslyChanged == 0))) break;
		previouslyChanged = currentlyChanged;

		currentBoard = checkBoard(currentBoard, newBoard);
		runs++;
	}

	int occupied = countOccupied(currentBoard);
	freeAllSeats(currentBoard);
	freeAllSeats(newBoard);
	return occupied;
}

int occupied2(char **array, int x, int y) {
	int columnMax = lineLength;
	int rowMax = arraySize;

	int occupiedCount = 0;
	for (int i = x+1; i < rowMax; i++) {
		if (array[i][y] == '#') {
			occupiedCount++;
			break;
		}

		if (array[i][y] != '.') break;
	}

	for (int i = x-1; i >= 0; i--) {
		if (array[i][y] == '#') {
			occupiedCount++;
			break;
		}

		if (array[i][y] != '.') break;
	}

	for (int i = y+1; i < columnMax; i++) {
		if (array[x][i] == '#') {
			occupiedCount++;
			break;
		}

		if (array[x][i] != '.') break;
	}

	for (int i = y-1; i >= 0; i--) {
		if (array[x][i] == '#') {
			occupiedCount++;
			break;
		}

		if (array[x][i] != '.') break;
	}

	int currentX = x - 1;
	int currentY = y - 1;
	while ((currentX >= 0) && (currentY >= 0)) {
		if (array[currentX][currentY] == '#') {
			occupiedCount++;
			break;
		}

		if (array[currentX][currentY] != '.') break;

		currentX--;
		currentY--;
	}

	currentX = x + 1;
	currentY = y + 1;
	while ((currentX < rowMax) && (currentY < columnMax)) {
		if (array[currentX][currentY] == '#') {
			occupiedCount++;
			break;
		}

		if (array[currentX][currentY] != '.') break;
		currentX++;
		currentY++;
	}

	currentX = x - 1;
	currentY = y + 1;
	while ((currentX >= 0) && (currentY < columnMax)) {
		if (array[currentX][currentY] == '#') {
			occupiedCount++;
			break;
		}

		if (array[currentX][currentY] != '.') break;
		currentX--;
		currentY++;
	}

	currentX = x + 1;
	currentY = y - 1;
	while ((currentX < rowMax) && (currentY >= 0)) {
		if (array[currentX][currentY] == '#') {
			occupiedCount++;
			break;
		}

		if (array[currentX][currentY] != '.') break;
		currentX++;
		currentY--;
	}

	return occupiedCount;
}

char **checkBoard2(char **currentBoard, char **newBoard) {
	currentlyChanged = 0;
	for (int i = 0; i < arraySize; i++) {
		char temp[lineLength+1];
		temp[lineLength] = 0;
		for (size_t j = 0; j < lineLength; j++) {
			char *current = &currentBoard[i][j];
			char new = *current;

			switch (*current) {
				case 'L':
					if (!occupied2(currentBoard, i, j)) {
						new = '#';
						currentlyChanged = 1;
					}
					break;
				case '#':
					if (occupied2(currentBoard, i, j) >= 5) {
						new = 'L';
						currentlyChanged = 1;
					}
					break;
			}

			temp[j] = new;
		}
		sprintf(newBoard[i], "%s", temp);
	}

	freeAllSeats(currentBoard);

	return newBoard;
}

int doRounds2(char **array) {
	char **currentBoard = array;
	int runs = 0;
	char **newBoard = NULL;
	while (1) {
		newBoard = allocateNewBoard(arraySize, lineLength+1);
		if (!newBoard) return -1;

		if (runs > 1 && ((currentlyChanged == 0) && (previouslyChanged == 0))) break;
		previouslyChanged = currentlyChanged;

		currentBoard = checkBoard2(currentBoard, newBoard);
		runs++;
	}

	int occupied = countOccupied(currentBoard);
	freeAllSeats(currentBoard);
	freeAllSeats(newBoard);
	return occupied;
}

char **copyBoard(char** array, int rows, int columns) {
	char **newBoard = malloc(sizeof(char*)*rows);
	if (!newBoard) return NULL;

	for (int i = 0; i < rows; i++) {
		newBoard[i] = malloc(columns);
		if (!newBoard[i]) return NULL;
		sprintf(newBoard[i], "%s", array[i]);
	}

	return newBoard;
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

	char **currentBoard = readSeats(fp);
	if (!currentBoard) {
		puts("Can't read numbers!");
		fclose(fp);
		return -1;
	}
	fclose(fp);

	lineLength = strlen(currentBoard[arraySize-1]);
	char **originalBoard = copyBoard(currentBoard, arraySize, lineLength+1);
	if (!originalBoard) return -1;

	printf("Part1: %d\n", doRounds(currentBoard));
	printf("Part2: %d\n", doRounds2(originalBoard));
}
