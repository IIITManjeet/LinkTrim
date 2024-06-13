import { useState } from 'react';
import './Form.css';
import { backendURL } from '../../constants';

export default function URLForm() {
    const [longUrl, setLongUrl] = useState('');
    const [alias, setAlias] = useState('');

    const handleSubmit = async (event: { preventDefault: () => void; }) => {
        event.preventDefault();
        try {
            const response = await fetch(backendURL + '/api/v1/shorts', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ target_url: longUrl, short_url: alias }),
            });
            if (response.ok) {
                console.log('Form submitted successfully!');
                setLongUrl('');
                setAlias('');
            } else {
                console.log('Form submission failed!');
            }
        } catch (error) {
            console.log('Error occurred during form submission:', error);
        }
    };
    return (
        <>
            <div className='virtual-body'>
                <div className="url-form">
                    <form onSubmit={handleSubmit}>
                        <div className='layer'>
                            <label htmlFor="long-url">Shorten a long URL:</label>
                            <input
                                type="text"
                                id="long-url"
                                placeholder="Enter a long URL"
                                value={longUrl}
                                onChange={(event) => setLongUrl(event.target.value)}
                                required />
                        </div>
                        <div className='layer'>
                            <label htmlFor="alias">Customize your link:</label>
                            <input type="text"
                                id="alias"
                                placeholder="Enter Alias (optional)"
                                value={alias}
                                onChange={(event) => setAlias(event.target.value)} />
                        </div>
                        <button type="submit" id='submit'>Shorten URL</button>
                    </form>
                </div>
            </div>
        </>
    )
}