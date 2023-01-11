#include <stdio.h>
#include <stdlib.h>

int reverse_captcha(int *captcha, int size, int offset)
{
    int answer = 0;
    for (int i = 0; i < size; i++)
    {
        if (captcha[i] == captcha[(i + offset) % size])
        {
            answer += captcha[i];
        }
    }
    return answer;
}

int main()
{
    int c;
    int size = 0; // 2125
    FILE *input = fopen("input.txt", "r");
    if (! input)
    {
        return 1;
    }
    do
    {
        c = fgetc(input);
        if (c == '\n')
        {
            break;
        }
        size++;
    } while(1);
    rewind(input);

    int captcha[size];
    for (int i = 0; i < size; i++)
    {
        c = fgetc(input);
        // ascii 48 equals 0, so if printf decimal of ascii c we will get
        // ascii 49, so if we do 49 - 48, we get the decimal 1.
        captcha[i] = (c - '0');
    }
    fclose(input);

    printf("next digit answer: %d\n", reverse_captcha(captcha, size, 1));
    printf("halfway around answer: %d\n", reverse_captcha(captcha, size, (size/2)));
}
