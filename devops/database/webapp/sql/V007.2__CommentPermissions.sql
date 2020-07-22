INSERT INTO permissions (human_name, description)
VALUES
    ('org.comments.list', 'Ability to view all comments in a thread. You must have view permissions on the parent resource.'),
    ('org.comments.create', 'Ability to post a new comment in a thread. You must have view permissions on the parent resource.'),
    ('org.comments.update', 'Ability to edit comments in a thread. You must have view permissions on the parent resource. Note that this grants the ability to edit all comments in the thread. You will always have the ability to edit your own comment.'),
    ('org.comments.delete', 'Ability to delete comments in a thread. You must have view permissions on the parent resource. Note that this grants the ability to delete all comments in the thread. You will always have the ability to delete your own comment.');
