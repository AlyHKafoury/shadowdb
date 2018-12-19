#include"statement_handler.h"
#include<iostream>
using namespace std;

PrepareResult prepare_statement(string &input_buffer, Statement &statement){
   if(input_buffer.substr(0,6) == "insert"){
       statement.type = STATEMENT_INSERT;
       return PREPARE_SUCCESS;
   } 
   if(input_buffer.substr(0,6) == "select"){
       statement.type = STATEMENT_SELECT;
       return PREPARE_SUCCESS;
   }
   return PREPARE_UNRECOGNIZED_STATEMENT;
}

void execute_statement(Statement &statement){
    switch(statement.type){
        case STATEMENT_INSERT:
            cout << "result of insert" << endl;
            break;
        case STATEMENT_SELECT:
            cout << "result of select" << endl;
            break;
    }
}