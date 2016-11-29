
#define QUEUE_SIZE 400
#define LEFT  1
#define UP    2
#define DOWN  3
#define RIGHT 4
#define MAX_X 24
#define MAX_Y 14
#define TILE_SIZE 32

typedef struct tag_node {
  char x;
  char y;
} node;

typedef struct tag_queue {
  int first;
  int last;
  int len;
  node elems[QUEUE_SIZE];
} queue;


void init();
void input();
int  update();
void render();
void pop_tail();
void push_head();
void gameover();
void draw_body();
void draw_head();
void draw_fruit();
void clear_tail();
void next_fruit();
void delay_10();
