describe :database do
  before do
    `rm db;touch db;`
  end

  def run_script(commands)
    raw_output = nil
    IO.popen('./shadowdb db', 'r+') do |pipe|
      commands.each do |command|
        pipe.puts command
      end

      pipe.close_write
      raw_output = pipe.gets(nil)
    end
    raw_output.split("\n")
  end

  it 'inserts and retreives a row' do
    result = run_script(['insert 1 aly email', 'select', '.exit'])
    expect(result).to match_array([
                                    'Shadow-DB *>> Added Row to table',
                                    'Executed command',
                                    'Shadow-DB *>> 1 aly email',
                                    'Executed command',
                                    'Shadow-DB *>> '
                                  ])
  end

  it 'prints error message when table is full' do
    script = (1..1401).map do |i|
      "insert #{i} user#{i} person#{i}@example.com"
    end
    script << '.exit'
    result = run_script(script)
    expect(result[-2]).to eq('Shadow-DB *>> Table is full')
  end

  it 'allows inserting strings that are the maximum length' do
    long_username = 'a' * 32
    long_email = 'a' * 255
    script = [
      "insert 1 #{long_username} #{long_email}",
      'select',
      '.exit'
    ]
    result = run_script(script)
    expect(result).to match_array([
                                    'Shadow-DB *>> Added Row to table',
                                    'Executed command',
                                    "Shadow-DB *>> 1 #{long_username} #{long_email}",
                                    'Executed command',
                                    'Shadow-DB *>> '
                                  ])
  end

  it 'prints error message if strings are too long' do
    long_username = 'a' * 33
    long_email = 'a' * 256
    script = [
      "insert 1 #{long_username} #{long_email}",
      'select',
      '.exit'
    ]
    result = run_script(script)
    expect(result).to match_array([
                                    'Shadow-DB *>> Input too long',
                                    'Shadow-DB *>> Executed command',
                                    'Shadow-DB *>> '
                                  ])
  end

  it 'keeps data after closing connection' do
    result1 = run_script([
                           'insert 1 user1 person1@example.com',
                           '.exit'
                         ])
    expect(result1).to match_array([
                                     'Executed command',
                                     'Shadow-DB *>> ',
                                     'Shadow-DB *>> Added Row to table'
                                   ])
    result2 = run_script([
                           'select',
                           '.exit'
                         ])
    expect(result2).to match_array([
                                     'Executed command',
                                     'Shadow-DB *>> ',
                                     'Shadow-DB *>> 1 user1 person1@example.com'
                                   ])
  end
end
