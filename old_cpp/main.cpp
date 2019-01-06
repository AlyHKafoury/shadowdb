#include<iostream>
#include"meta_handler.h"
#include"statement_handler.h"
using namespace std;

void print_prompt(){
    cout << "Shadow-DB >>> ";
}

int main(int argc, char* argv[]) {
    string input_buffer;
    while(true){
        print_prompt();
        getline(cin, input_buffer);
        if(input_buffer[0] == *".") {
            switch (do_meta_command(input_buffer)){
                case (META_COMMAND_SUCCESS):
                    continue;
                case (META_COMMAND_UNRECOGNIZED_COMMAND):
                    cout << "Unknown Command : " << input_buffer << endl;
                    continue;
            }
        }
        Statement statement;
        switch (prepare_statement(input_buffer, statement)) {
            case PREPARE_SUCCESS:
                execute_statement(statement);
                break;
            case PREPARE_UNRECOGNIZED_STATEMENT:
                cout << "Unrecognized Keyword at the start of command" << endl;
                continue;
        }
        cout << "Executed command !" << endl;
    }
    

}