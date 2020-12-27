#include <iostream>

int main(int argc, char *argv[])
{
    std::cout << "hello clang(++)" << std::endl;

#ifdef _WIN32
    system("pause");
#endif // !_WIN32
    return 0;
}