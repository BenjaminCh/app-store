export default function AppStoreApp ({ children }) {
    return (
        <div class="container">
            <div class="grid">
                <div class="col-12">
                    <h1> App Store </h1>
                    <p> Discover plenty wonderful apps ! Just type what you have in mind. </p>
                    { children }
                </div>
            </div>
        </div>
    );
}