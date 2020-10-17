module NoParameters
end

World NoParameters

Given('no parameters') do
  @paramcount = 0 # Write code here that turns the phrase above into concrete actions
end

When('I count my parameters') do
  @actualcount = ARGV.length # Write code here that turns the phrase above into concrete actions
end

Then('I should see {int}') do |expectedcount|
  expect(@actualcount).to eq(expectedcount) # Write code here that turns the phrase above into concrete actions
end
