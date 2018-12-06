#!/bin/python
from datetime import datetime, timedelta, time

nb_hours = timedelta(0, 0, 0, 0, 40, 7, 0)

now_str = raw_input('At what time did you arrived (HH:MM): ')
now = datetime.strptime(now_str, '%H:%M')

try:
    lunch_int = input('How many minutes for the lunch: ')
except ValueError:
    print 'Please, enter a number...'
    raise ValueError
hours_lunch = timedelta(0, 0, 0, 0, lunch_int, 0, 0)

print 'You shall work for {}'.format(nb_hours)
print 'Given that, you can leave at: {}'.format(datetime.strftime(nb_hours + now + hours_lunch, '%H:%M'))
