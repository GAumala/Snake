#include "snake.h"

int main(void)
{
    init();
    render();
    for (;;) {
        input();
        if (update()) {
            gameover();
        }
        render();
        delay_10();
    }
    return 0;
}

