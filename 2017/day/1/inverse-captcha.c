#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int reverse_captcha(char *captcha, int len, int offset)
{
    int answer = 0;
    for (int i = 0; i < len; i++)
    {
        if (captcha[i] == captcha[(i + offset) % len])
        {
            answer += captcha[i] - '0';
        }
    }

    return answer;
}

int main()
{
    FILE *stream;
    stream = fopen("input.txt", "r");
    if (stream == NULL) {
        perror("fopen");
        exit(EXIT_FAILURE);
    }

    char *line = NULL;
    size_t len = 0;
    while (getline(&line, &len, stream) != -1) {
        line[strcspn(line, "\n")] = 0;
        len = strlen(line);
        printf("next digit answer: %d\n", reverse_captcha(line, len, 1));
        printf("halfway around answer: %d\n", reverse_captcha(line, len, (len/2)));
    }

    free(line);
    fclose(stream);
    exit(EXIT_SUCCESS);
}
