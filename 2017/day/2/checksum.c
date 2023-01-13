#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int checksum(int *numbers, int size)
{
        int min = INT_MAX;
        int max = INT_MIN;
        for (int i = 0; i < size; i++) {
            if (numbers[i] > max)
                max = numbers[i];
            if (numbers[i] < min)
                min = numbers[i];
        }

        return (max - min);
}

int divisible(int *numbers, int size)
{
    int number1;
    int number2;
    int divisor;

    for (int i = 0; i < size; i++) {
        number1 = numbers[i];
        for (int j = 0; j < size; j++) {
            number2 = numbers[j];
            if (number1 > number2) {
                if (number1 % number2 == 0 || number2 % number1 == 0) {
                    return (number1 / number2);
                }
            }
        }
    }
}

int main()
{
    FILE *stream;
    stream = fopen("input.txt", "r");
    if (stream == NULL) {
        perror("fopen");
        exit(EXIT_FAILURE);
    }

    int size = 0;
    int sum = 0;
    int divisor = 0;
    char *line = NULL;
    size_t len = 0;

    while (getline(&line, &len, stream) != -1) {
        int numbers[16];
        int i = 0;
        char *number = strtok(line, "\t");

        while (number != NULL) {
            numbers[i++] = atoi(number);
            number = strtok(NULL, "\t");
        }

        size = (sizeof(numbers)/sizeof(int));
        sum += checksum(numbers, size);
        divisor += divisible(numbers, size);
    }

    printf("checksum answer: %d, divisible answer: %d\n", sum, divisor);

    free(line);
    fclose(stream);
    exit(EXIT_SUCCESS);
}
