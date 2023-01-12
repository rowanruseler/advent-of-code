#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/*
 * Function: checksum
 * ------------------
 *  iterates over the array to determine the difference between the largest
 *  value and the smallest value.
 *
 *  *numbers: a pointer to an array of numbers
 *  size: the size of the array
 *
 *  returns: the difference between the largest numbers and smallest numbers
 *           by substracting the smallest number over the largest number.
 *
 */
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

/*
 * Function: divisible
 * -------------------
 *  iterates twice over the array to find the only two numbers where one
 *  evenly divides the other - that is, where the result of the division
 *  operation is a whole number.
 *
 *  *numbers: a pointer to an array of numbers
 *  size: the size of the array
 *
 *  returns: the division of the divisors by dividing the two numbers that
 *           are divisors
 */
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
                    divisor = (number1 / number2);
                }
            }
        }
    }

    return (divisor);
}

int main()
{
    FILE *stream;
    stream = fopen("input.txt", "r");
    if (stream == NULL) {
        perror("fopen");
        exit(EXIT_FAILURE); // ref: man 3 exit
    }

    int size = 0;
    int sum = 0;
    int divisor = 0;
    char *line = NULL;
    size_t len = 0; // ref: man 3 size_t

    // ref: man 3 getline
    while (getline(&line, &len, stream) != -1) {
        int numbers[16];
        int i = 0;
        char *number = strtok(line, "\t"); // ref: man 3 strtok

        while (number != NULL) {
            numbers[i++] = atoi(number); // ref: man 3 atoi
            number = strtok(NULL, "\t");
        }

        size = (sizeof(numbers)/sizeof(int));
        sum += checksum(numbers, size);
        divisor += divisible(numbers, size);
    }

    printf("checksum answer: %d, divisible answer: %d\n", sum, divisor);

    free(line); // ref: man 3 free
    fclose(stream);
    exit(EXIT_SUCCESS);
}
