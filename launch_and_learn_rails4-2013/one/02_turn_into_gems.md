!SLIDE
# Turn into gems #

!SLIDE smaller
# activerecord-deprecated_finders
    @@@ ruby
    # find_all_by_
    Rails 3: Ticket.find_all_by_name('ticket 1')
    Rails 4: Ticket.where(name: 'ticket 1')

    # find_last_by_
    Rails 3: Ticket.find_last_by_name('ticket 1')
    Rails 4: Ticket.where(name: 'ticket 1').last

    # find_or_create_by_
    Rails 3: Ticket.find_or_create_by_name('ticket 1')
    Rails 4: Ticket.where(name: 'ticket 1').first_or_create

!SLIDE smaller
# activerecord-deprecated_finders
    @@@ ruby
    Rails 3: Ticket.find(:all, conditions: { name: 'ticket 1' })
    Rails 4: Ticket.where(name: 'ticket 1')

!SLIDE
# Page and action caching

!SLIDE bullets incremental
# rails-observers

* Active Record Observer
* Action Controller Sweeper

!SLIDE
# activerecord-session_store

!SLIDE
# activeresource
