#include <limits.h>
#include <stdio.h>

int is_digit(int c)
{
    return c >= '0' && c <= '9';
}

int main()
{
    FILE *input = fopen("input.txt", "r");
    int c, sum, min = INT_MAX, max = INT_MIN, current = 0;
    int i = 0;
    if (! input)
    {
        return 1;
    }
    do
    {
        c = fgetc(input);
        if (! is_digit(c))
        {
            if (current < min)
                min = current;
            if (current > max)
                max = current;
            current = 0;
        }
        if (c == '\n')
        {
            sum += max - min;
            printf("line: %d, min: %d\n", i, min);
            printf("line: %d, max: %d\n", i, max);
            min = INT_MAX;
            max = INT_MIN;
            printf("%d\n", sum);
            i++;
        }
        else if (c == EOF)
            break;
        else if (is_digit(c))
            current = current * 10 + (c - '0');
    } while(1);

    fclose(input);
}
