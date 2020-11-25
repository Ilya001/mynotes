box.cfg{listen = 3301}
notes = box.schema.space.create('notes')
notes:format({
    {name='id', type='unsigned'},
    {name='text', type='string'}
})
notes:create_index('primary', {
    type='tree',
    parts={'id'}
})
notes:create_index('scanner', {
    type='tree',
    parts={'id', 'text'}
})
notes:auto_increment({'<div contenteditable="true">Новая заметка...</div>'})

function auto_increment_text(text)
    data = notes:auto_increment({text})
    return data
end

function update_text(id, text)
    data = notes:update({id}, {{'=', 2, text}})
    return data
end