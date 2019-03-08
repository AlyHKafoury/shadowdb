describe :database do
  def run_script(commands)
    raw_output = nil
    IO.popen('./shadowdb', 'r+') do |pipe|
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
end
