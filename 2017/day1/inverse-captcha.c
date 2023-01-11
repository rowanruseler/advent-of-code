#include <stdio.h>

void reverse_captcha(int *captcha, int size) {
    int answer = 0;
    for (int x = 0; x < size; x++) {
        if (captcha[x] == captcha[(x + 1) % size]) {
            answer += captcha[x];
        }
    }
    printf("answer: %d\n", answer);
}

int main() {
    int captcha[] = {1,1,2,2};
    int size = (sizeof(captcha)/sizeof(int));
    reverse_captcha(captcha, size);
}
